<template>
  <footer class="d-flex align-items-center justify-content-center mt-5">
    Copyright (c) {{ new Date().getFullYear() }} Thomas. All Rights Reserved.
  </footer>
</template>

<style scoped>
footer {
  background-color: var(--bs-body-color);
  color: white;
  min-height: 3rem;
  font-size: 0.8rem;
  width: 100%;
  position: fixed;
  bottom: 0;
}
</style>
