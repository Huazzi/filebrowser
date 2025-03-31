<template>
  <footer class="icp-container">
    <div class="icp-wrapper">
      <span class="icp-copyright">© {{ year }} {{ author }}</span>
      <a
        class="icp-link"
        href="http://beian.miit.gov.cn/"
        target="_blank"
        rel="noopener noreferrer"
      >
        <span class="icp-icon"></span>
        {{ record }}
      </a>
    </div>
  </footer>
</template>

<script setup>
const year = new Date().getFullYear();
const author = "huazzi";
const record = "黔ICP备2025xxxxxx号";
</script>

<style scoped>
/* 自定义字体声明 */
@font-face {
  font-family: "ShangTuDongGuan";
  src: url("@/assets/fonts/custom/上图东观体-常规.woff2") format("woff2");
  font-weight: normal;
  font-style: normal;
  /* 确保字体加载时的平滑过渡 */
  font-display: swap;
}
.icp-container {
  width: 100%;
  background: #2d2d2d;
  box-shadow: 0 -1px 6px rgba(0, 0, 0, 0.3);
  font-family:
    "ShangTuDongGuan",
    -apple-system,
    BlinkMacSystemFont,
    "Segoe UI",
    Roboto,
    Oxygen,
    Ubuntu,
    Cantarell,
    "Open Sans",
    "Helvetica Neue",
    sans-serif;
  margin-top: auto; /* 确保footer停留在底部 */
}

.icp-wrapper {
  max-width: 1200px;
  margin: 0 auto;
  padding: 12px 20px;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  font-size: 14px;
  line-height: 1.5;
  color: rgba(255, 255, 255, 0.8);
}

.icp-copyright {
  position: relative;
  padding-right: 10px;
}

.icp-copyright::after {
  content: "•";
  position: absolute;
  right: 0;
  color: rgba(255, 255, 255, 0.3);
}

.icp-link {
  display: inline-flex;
  align-items: center;
  color: rgba(255, 255, 255, 0.8);
  text-decoration: none;
  transition: color 0.2s ease;
}

.icp-link:hover {
  color: #7eb6f0;
}

.icp-icon {
  display: inline-block;
  width: 14px;
  height: 14px;
  margin-right: 6px;
  background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 24 24' fill='none' stroke='rgba(255, 255, 255, 0.8)' stroke-width='2' stroke-linecap='round' stroke-linejoin='round'%3E%3Cpath d='M22 12h-4l-3 9L9 3l-3 9H2'/%3E%3C/svg%3E");
  background-position: center;
  background-repeat: no-repeat;
  background-size: contain;
  opacity: 0.7;
  transition: opacity 0.2s ease;
}

.icp-link:hover .icp-icon {
  opacity: 1;
  background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 24 24' fill='none' stroke='%237eb6f0' stroke-width='2' stroke-linecap='round' stroke-linejoin='round'%3E%3Cpath d='M22 12h-4l-3 9L9 3l-3 9H2'/%3E%3C/svg%3E");
}

@media (max-width: 768px) {
  .icp-wrapper {
    flex-direction: column;
    padding: 8px 16px;
  }

  .icp-copyright {
    padding-right: 0;
  }

  .icp-copyright::after {
    display: none;
  }
}
</style>
