<script setup lang="ts">
import AppForm from './forms/AppForm.vue'
import TextInput from './forms/TextInput.vue'
</script>

<template>
  <div class="container">
    <div class="row">
      <div class="col">
        <h1 class="mt-5">Login</h1>
        <hr />
        <AppForm method="POST" action="/login">
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
        </AppForm>
      </div>
    </div>
  </div>
</template>
