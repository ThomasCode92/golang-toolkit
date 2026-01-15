<script setup lang="ts">
import { ref } from 'vue'

import AppForm from './forms/AppForm.vue'
import TextInput from './forms/TextInput.vue'

const email = ref('')
const password = ref('')

const handleSubmit = (formData: FormData) => {
  console.log('Form submitted successfully')
  console.log('Email:', formData.get('email'))
  console.log('Password:', formData.get('password'))

  // TODO: Implement actual login logic here
  // Example: await loginUser(formData.get('email'), formData.get('password'))
}
</script>

<template>
  <div class="container">
    <div class="row">
      <div class="col">
        <h1 class="mt-5">Login</h1>
        <hr />
        <AppForm method="POST" action="/login" @onSubmit="handleSubmit">
          <text-input
            v-model="email"
            label="Email"
            type="email"
            name="email"
            placeholder="Enter your email"
            required
          />
          <text-input
            v-model="password"
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
