document.getElementById("submit").addEventListener("click", function(event) {
    let isValid = true;

    // Clear previous errors
    document.querySelectorAll('.error-message').forEach(e => e.textContent = '');
    document.querySelectorAll('input, select, textarea').forEach(e => e.classList.remove('input-error'));

    // Password Validation
    const password = document.getElementById("password");
    if (!password.value.trim()) {
        displayError(password, "Please input a password");
        isValid = false;
    }

    // Email Address Validation
    const emailAddress = document.getElementById("email_address");
    if (!emailAddress.value.trim() || !/\S+@\S+\.\S+/.test(emailAddress.value)) {
        displayError(emailAddress, "A valid email address is required");
        isValid = false;
    }

    if (!isValid) {
        event.preventDefault();
        return; // Stop the function if the form is not valid
    }

    if(isValid) {
        // Only proceed if the form is valid
        const formData = {
            email_address: document.getElementById("email_address").value,
            password: document.getElementById("password").value,
        };

        fetch('/lawyer_login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(formData),
        })
        .then(response => response.json())
        .then(data => {
            
            if (data.error) {
                document.getElementById("unknownAccountMessage").style.display = "block";
                console.log(data.error);
            }
            else {
                console.log('Success:', data);
                lawyer_id = data.lawyer_id;
                console.log(lawyer_id); 
                window.location.href = "/display-cases";
            }
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
