<script setup lang="ts">
defineProps(['label', 'name', 'items', 'required'])
</script>

<template>
  <div class="mb-3">
    <label :for="name" class="form-label">{{ label }}</label>
    <select :id="name" class="form-select" :name="name" :required="required">
      <option v-for="item in items" :key="item.id" :value="item.value">
        {{ item.text }}
      </option>
    </select>
  </div>
</template>
