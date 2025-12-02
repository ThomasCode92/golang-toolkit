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
  `,
};

const SelectInput = {
  props: ["items", "name", "label", "required"],
  template: `
    <div class="mb-3">
      <label :for="name" class="form-label">{{label}}</label>
      <select :id="name" class="form-select" :name="name" :required="required">
        <option v-for="item in items" :key="item.id" :value="item.value">
          {{ item.text }}
        </option>
      </select>
    </div>
  `,
};

const CheckboxInput = {
  props: ["name", "label", "required", "value"],
  template: `
    <div class="form-check mb-3">
      <input
        :id="name"
        type="checkbox"
        :name="name"
        class="form-check-input"
        :value="value"
        :required="required"
      />
      <label :for="name" class="form-check-label">{{ label }}</label>
    </div>
  `,
};
