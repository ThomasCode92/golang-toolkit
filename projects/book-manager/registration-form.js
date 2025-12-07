const RegistrationForm = {
  data() {
    return {
      addressSameChecked: true,
    };
  },
  props: ["items"],
  template: `
    <h3>Registration</h3>
    <hr />
    <form class="needs-validation" autocomplete="off" novalidate>
      <text-input label="First Name" type="text" name="first-name" required></text-input>
      <text-input label="Last Name" type="text" name="last-name" required></text-input>
      <text-input label="Email" type="email" name="email" required></text-input>
      <text-input label="Password" type="password" name="password" required></text-input>
      <select-input label="Favorite Color" name="favorite-color" :items="items" required ></select-input>
      <checkbox-input label="I agree to the Terms and Conditions" required></checkbox-input>

      <text-input label="Address" type="text" name="address" required="true"></text-input>
      <text-input label="City/Town" type="text" name="city" required="true"></text-input>
      <text-input label="State/Province" type="text" name="province" required="true"></text-input>
      <text-input label="Zip/Postal" type="text" name="zip" required="true"></text-input>

      <checkbox-input v-on:click="addressSame" label="Mailing Address Same" checked="true" v-model="addressSameChecked"></checkbox-input>

      <div v-if="addressSameChecked === false">
        <div class="mt-3">
          <text-input label="Mailing Address" type="text" name="address2"></text-input>
          <text-input label="Mailing City/Town" type="text" name="city2"></text-input>
          <text-input label="Mailing State/Province" type="text" name="province2"></text-input>
          <text-input label="Mailing Zip/Postal" type="text" name="zip2"></text-input>
        </div>
      </div>

      <hr />
      <input type="submit" class="btn btn-outline-primary" value="Register"/>
    </form>
  `,
  methods: {
    addressSame() {
      this.addressSameChecked = !this.addressSameChecked;
    },
  },
  mounted() {
    (() => {
      "use strict";

      // Fetch all the forms we want to apply custom Bootstrap validation styles to
      const forms = document.querySelectorAll(".needs-validation");

      // Loop over them and prevent submission
      Array.from(forms).forEach((form) => {
        form.addEventListener(
          "submit",
          (event) => {
            if (!form.checkValidity()) {
              event.preventDefault();
              event.stopPropagation();
            }

            form.classList.add("was-validated");
          },
          false,
        );
      });
    })();
  },
  components: {
    "text-input": TextInput,
    "select-input": SelectInput,
    "checkbox-input": CheckboxInput,
  },
};
