<script setup lang="ts">
import { onMounted } from 'vue'

import { validateForms } from '@/utils/validation'

defineProps(['method', 'action'])

const emit = defineEmits<{
  onSubmit: [formData: FormData]
}>()

onMounted(() => {
  validateForms()
})

const onSubmit = (event: Event) => {
  const form = event.target as HTMLFormElement
  const isValid = form.checkValidity() // NOTE: Bootstrap validation utility

  if (!isValid) return

  const formData = new FormData(form)
  emit('onSubmit', formData)

  // NOTE: Bootstrap validation will handle styling if form is invalid,
  // so we don't need to duplicate that logic here
}
</script>

<template>
  <form
    :method="method"
    :action="action"
    @submit.prevent="onSubmit"
    class="needs-validation"
    novalidate
    autocomplete="off"
  >
    <slot />
  </form>
</template>
