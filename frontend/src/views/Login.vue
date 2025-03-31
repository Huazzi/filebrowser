<template>
  <div id="login" :class="{ recaptcha: recaptcha }">
    <div class="login-container">
      <form @submit="submit" class="login-form">
        <div class="logo-container">
          <img :src="logoURL" alt="File Browser" class="logo" />
          <h1 class="app-name">{{ name }}</h1>
        </div>
        
        <div v-if="error !== ''" class="error-message">
          <span class="error-icon">!</span>
          <span>{{ error }}</span>
        </div>

        <div class="input-group">
          <label for="username">{{ t('login.username') }}</label>
          <input
            id="username"
            autofocus
            class="form-input"
            type="text"
            autocapitalize="off"
            v-model="username"
            :placeholder="t('login.username')"
          />
        </div>
        
        <div class="input-group">
          <label for="password">{{ t('login.password') }}</label>
          <input
            id="password"
            class="form-input"
            type="password"
            v-model="password"
            :placeholder="t('login.password')"
          />
        </div>
        
        <div v-if="createMode" class="input-group">
          <label for="passwordConfirm">{{ t('login.passwordConfirm') }}</label>
          <input
            id="passwordConfirm"
            class="form-input"
            type="password"
            v-model="passwordConfirm"
            :placeholder="t('login.passwordConfirm')"
          />
        </div>

        <div v-if="recaptcha" id="recaptcha" class="recaptcha-container"></div>
        
        <button
          class="submit-button"
          type="submit"
        >
          {{ createMode ? t('login.signup') : t('login.submit') }}
        </button>

        <p class="toggle-mode" @click="toggleMode" v-if="signup">
          {{ createMode ? t("login.loginInstead") : t("login.createAnAccount") }}
        </p>
      </form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { StatusError } from "@/api/utils";
import * as auth from "@/utils/auth";
import {
  name,
  logoURL,
  recaptcha,
  recaptchaKey,
  signup,
} from "@/utils/constants";
import { inject, onMounted, ref } from "vue";
import { useI18n } from "vue-i18n";
import { useRoute, useRouter } from "vue-router";

// Define refs
const createMode = ref<boolean>(false);
const error = ref<string>("");
const username = ref<string>("");
const password = ref<string>("");
const passwordConfirm = ref<string>("");

const route = useRoute();
const router = useRouter();
const { t } = useI18n({});
// Define functions
const toggleMode = () => (createMode.value = !createMode.value);

const $showError = inject<IToastError>("$showError")!;

const submit = async (event: Event) => {
  event.preventDefault();
  event.stopPropagation();

  const redirect = (route.query.redirect || "/files/") as string;

  let captcha = "";
  if (recaptcha) {
    captcha = window.grecaptcha.getResponse();

    if (captcha === "") {
      error.value = t("login.wrongCredentials");
      return;
    }
  }

  if (createMode.value) {
    if (password.value !== passwordConfirm.value) {
      error.value = t("login.passwordsDontMatch");
      return;
    }
  }

  try {
    if (createMode.value) {
      await auth.signup(username.value, password.value);
    }

    await auth.login(username.value, password.value, captcha);
    router.push({ path: redirect });
  } catch (e: any) {
    // console.error(e);
    if (e instanceof StatusError) {
      if (e.status === 409) {
        error.value = t("login.usernameTaken");
      } else if (e.status === 403) {
        error.value = t("login.wrongCredentials");
      } else {
        $showError(e);
      }
    }
  }
};

// Run hooks
onMounted(() => {
  if (!recaptcha) return;

  window.grecaptcha.ready(function () {
    window.grecaptcha.render("recaptcha", {
      sitekey: recaptchaKey,
    });
  });
});
</script>

<style scoped>
@import url("@/css/login.css");
</style>