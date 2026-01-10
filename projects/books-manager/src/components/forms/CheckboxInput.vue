<script setup lang="ts">
defineProps(['label', 'name', 'value', 'checked', 'required'])
</script>

<template>
  <div class="form-check mb-3">
    <input
      :id="name"
      type="checkbox"
      :name="name"
      class="form-check-input"
      :value="value"
      :checked="checked"
      :required="required"
    />
    <label :for="name" class="form-check-label">{{ label }}</label>
  </div>
</template>
