const RegistrationForm = {
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
      <hr />
      <input type="submit" class="btn btn-outline-primary" value="Register"/>
    </form>
  `,
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
