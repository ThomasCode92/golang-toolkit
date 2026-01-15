<script setup lang="ts">
defineProps({
  name: { type: String },
  label: { type: String },
  type: { type: String },
  placeholder: { type: String },
  value: { type: String },
  required: { type: Boolean },
  min: { type: String },
  max: { type: String },
})
</script>

<template>
  <div class="mb-3">
    <label :for="name" class="form-label">{{ label }}</label>
    <input
      :id="name"
      :type="type"
      :name="name"
      class="form-control"
      :placeholder="placeholder"
      :autocomplete="name + '-new'"
      :value="value"
      :required="required"
      :min="min"
      :max="max"
    />
  </div>
</template>
