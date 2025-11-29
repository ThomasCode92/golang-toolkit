const TextInput = {
  props: {
    name: String,
    type: String,
    label: String,
    placeholder: String,
    required: String,
    min: String,
    max: String,
    value: String,
  },
  template: `
    <div class="mb-3">
      <label :for="name" class="form-label">{{ label }}</label>
      <input
        :id="name"
        :type="type"
        class="form-control"
        :placeholder="placeholder"
        :autocomplete="name + '-new'"
        :value="value"
        :required="required"
        :min="min"
        :max="max"
      />
    </div>
  `,
};
