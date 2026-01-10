<script setup lang="ts">
import { onMounted } from 'vue'

import TextInput from './forms/TextInput.vue'
import { validateForms } from '@/utils/validation'

onMounted(() => {
  validateForms()
})
</script>

<template>
  <div class="container">
    <div class="row">
      <div class="col">
        <h1 class="mt-5">Login</h1>
        <hr />
        <form method="POST" action="/login" class="needs-validation" novalidate>
          <text-input
            label="Email"
            type="email"
            name="email"
            placeholder="Enter your email"
            required
          />
          <text-input
            label="Password"
            type="password"
            name="password"
            placeholder="Enter your password"
            required
          />
          <hr />
          <input type="submit" class="btn btn-primary" value="Login" />
        </form>
      </div>
    </div>
  </div>
</template>
