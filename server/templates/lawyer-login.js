// Javascript for the lawyer login page
// Form validation and form submission logic for the login functionality
// Original for this project with ChatGPT assistance

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
                securitystring = data.securitystring;
                timestamp = data.timestamp;
                console.log(lawyer_id); 
                
                //Three cookies to verify security. We could store these all in a JSON in one cookie, but that would 
                //require further encoding since the JSON was breaking the cookie format. 
                //We can't use HTTPS on localhost but in an actual implementation, we would use HTTPS. If we w

                const tomorrow = new Date();
                tomorrow.setDate(tomorrow.getDate() + 1);
                document.cookie = "lawyer_id=" + lawyer_id + "; SameSite=Strict; expires=" + tomorrow.toUTCString();
                document.cookie = "securitystring=" + securitystring + "; SameSite=Strict; expires=" + tomorrow.toUTCString();
                document.cookie = "securitytimestamp=" + timestamp + "; SameSite=Strict; expires=" + tomorrow.toUTCString();
    

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
