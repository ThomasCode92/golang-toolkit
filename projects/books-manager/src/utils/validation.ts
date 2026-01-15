//**
// * Validate Forms using Bootstrap's custom validation styles
// * https://getbootstrap.com/docs/5.3/forms/validation/#custom-styles
// */
export const validateForms = () => {
  'use strict'

  // Fetch all the forms we want to apply custom Bootstrap validation styles to
  const forms = document.querySelectorAll('.needs-validation') as NodeListOf<HTMLFormElement>

  // Loop over them and prevent submission
  Array.from(forms).forEach((form) => {
    form.addEventListener(
      'submit',
      (event) => {
        if (!form.checkValidity()) {
          event.preventDefault()
          event.stopPropagation()
        }

        form.classList.add('was-validated')
      },
      false,
    )
  })
}
