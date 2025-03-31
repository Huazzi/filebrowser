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
          <input id="username" autofocus class="form-input" type="text" autocapitalize="off" v-model="username"
            :placeholder="t('login.username')" />
        </div>

        <div class="input-group">
          <label for="password">{{ t('login.password') }}</label>
          <input id="password" class="form-input" type="password" v-model="password"
            :placeholder="t('login.password')" />
        </div>

        <div v-if="createMode" class="input-group">
          <label for="passwordConfirm">{{ t('login.passwordConfirm') }}</label>
          <input id="passwordConfirm" class="form-input" type="password" v-model="passwordConfirm"
            :placeholder="t('login.passwordConfirm')" />
        </div>

        <div v-if="recaptcha" id="recaptcha" class="recaptcha-container"></div>

        <button class="submit-button" type="submit">
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
:root {
  --primary-color: #3498db;
  --primary-dark: #2980b9;
  --error-color: #e74c3c;
  --success-color: #2ecc71;
  --bg-color: #f5f7fa;
  --card-color: #ffffff;
  --text-color: #2c3e50;
  --border-color: #e1e4e8;
  --shadow-color: rgba(0, 0, 0, 0.1);
}

#login {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  background-color: var(--bg-color);
  padding: 20px;
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Oxygen, Ubuntu, Cantarell, "Open Sans", "Helvetica Neue", sans-serif;
}

.login-container {
  width: 100%;
  max-width: 420px;
  padding: 2rem;
  background-color: var(--card-color);
  border-radius: 12px;
  box-shadow: 0 8px 24px var(--shadow-color);
}

.logo-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-bottom: 2rem;
}

.logo {
  max-width: 80px;
  height: auto;
  margin-bottom: 1rem;
}

.app-name {
  margin: 0;
  color: var(--text-color);
  font-size: 1.8rem;
  font-weight: 600;
}

.error-message {
  display: flex;
  align-items: center;
  padding: 12px;
  margin-bottom: 1.5rem;
  background-color: rgba(231, 76, 60, 0.1);
  border-left: 4px solid var(--error-color);
  border-radius: 4px;
  color: var(--error-color);
}

.error-icon {
  display: inline-flex;
  justify-content: center;
  align-items: center;
  width: 24px;
  height: 24px;
  background-color: var(--error-color);
  color: white;
  border-radius: 50%;
  margin-right: 10px;
  font-weight: bold;
}

.input-group {
  margin-bottom: 1.5rem;
}

.input-group label {
  display: block;
  margin-bottom: 8px;
  font-weight: 500;
  color: var(--text-color);
  font-size: 0.9rem;
}

.form-input {
  width: 100%;
  padding: 12px 16px;
  font-size: 1rem;
  border: 2px solid var(--border-color);
  border-radius: 8px;
  background-color: white;
  transition: border-color 0.3s, box-shadow 0.3s;
}

.form-input:focus {
  outline: none;
  border-color: var(--primary-color);
  box-shadow: 0 0 0 3px rgba(52, 152, 219, 0.25);
}

.form-input::placeholder {
  color: #a0aec0;
}

.recaptcha-container {
  margin: 1.5rem 0;
}

.submit-button {
  width: 100%;
  padding: 14px;
  background-color: var(--primary-color);
  color: white;
  border: none;
  border-radius: 8px;
  font-size: 1rem;
  font-weight: 600;
  cursor: pointer;
  transition: background-color 0.3s;
}

.submit-button:hover {
  background-color: var(--primary-dark);
}

.toggle-mode {
  margin-top: 1.5rem;
  text-align: center;
  color: var(--primary-color);
  cursor: pointer;
  font-size: 0.9rem;
}

.toggle-mode:hover {
  text-decoration: underline;
}

@media (max-width: 480px) {
  .login-container {
    padding: 1.5rem;
  }
}
</style>