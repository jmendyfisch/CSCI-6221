document.getElementById("startIntakeBtn").addEventListener("click", function() {
    document.getElementById("intakeForm").style.display = "block";
    this.style.display = "none";
});




document.getElementById("submit").addEventListener("click", function(event) {
    let isValid = true;

    // Clear previous errors
    document.querySelectorAll('.error-message').forEach(e => e.textContent = '');
    document.querySelectorAll('input, select, textarea').forEach(e => e.classList.remove('input-error'));

    // First Name Validation
    const firstName = document.getElementById("client_first_name");
    if (!firstName.value.trim()) {
        displayError(firstName, "First Name is required");
        isValid = false;
    }

    // Last Name Validation
    const lastName = document.getElementById("client_last_name");
    if (!lastName.value.trim()) {
        displayError(lastName, "Last Name is required");
        isValid = false;
    }

    // Case Type Validation
    const caseType = document.getElementById("type");
    if (caseType.value === "Please Select") {
        displayError(caseType, "Please select a case type");
        isValid = false;
    }

    // Description Validation
    const description = document.getElementById("description");
    if (!description.value.trim()) {
        displayError(description, "A description is required");
        isValid = false;
    }


    const phoneNumber = document.getElementById("phone_number");
    const phoneNumberRegex = /^\d{10}$/;
    if (!phoneNumberRegex.test(phoneNumber.value)) {
        displayError(phoneNumber, "Phone number must be 10 digits");
        isValid = false;
    }

    // Email Address Validation
    const emailAddress = document.getElementById("email_address");
    if (!emailAddress.value.trim() || !/\S+@\S+\.\S+/.test(emailAddress.value)) {
        displayError(emailAddress, "A valid email address is required");
        isValid = false;
    }

    //alert("Form is valid: " + isValid);

    if (!isValid) {
        event.preventDefault();
        return; // Stop the function if the form is not valid
    }

    if(isValid) {
        // Only proceed if the form is valid
        const formData = {
            client_first_name: document.getElementById("client_first_name").value,
            client_last_name: document.getElementById("client_last_name").value,
            type: document.getElementById("type").value,
            description: document.getElementById("description").value,
            phone_number: document.getElementById("phone_number").value,
            email_address: document.getElementById("email_address").value
        };

        fetch('/create_case', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(formData),
        })
        .then(response => response.json())
        .then(data => {
            console.log('Success:', data);
            document.getElementById("intakeForm").style.display = "none";
            document.getElementById("thankYouMessage").style.display = "block";
        })
        .catch((error) => {
            console.error('Error:', error);
        });
    }
});

function displayError(element, message) {
    const errorMessage = document.getElementById(element.id + "_error") 
    errorMessage.textContent = message;
    element.classList.add('input-error');
}
