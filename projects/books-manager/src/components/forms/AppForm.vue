<script setup lang="ts">
import { onMounted } from 'vue'
import { validateForms } from '@/utils/validation'

defineProps(['method', 'action'])

onMounted(() => {
  validateForms()
})
</script>

<template>
  <form :method="method" :action="action" class="needs-validation" novalidate autocomplete="off">
    <slot />
  </form>
</template>
