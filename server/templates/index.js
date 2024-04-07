
document.getElementById("startIntakeBtn").addEventListener("click", function() {
    document.getElementById("intakeForm").style.display = "block";
    this.style.display = "none";
});

document.getElementById("intakeForm").addEventListener("submit", function(event) {
    event.preventDefault();

    const formData = {
        client_first_name: document.getElementById("client_first_name").value,
        client_last_name: document.getElementById("client_last_name").value,
        type: document.getElementById("type").value,
        description: document.getElementById("description").value,
        phone_number: document.getElementById("phone_number").value,
        email_address: document.getElementById("email_address").value
    };

    fetch(this.action, {
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
});


document.addEventListener("DOMContentLoaded", function() {
    const intakeForm = document.getElementById("intakeForm");
    intakeForm.addEventListener("submit", function(event) {
        let hasErrors = false;
        
        // Validate First Name
        if (!document.getElementById("client_first_name").value.trim()) {
            alert("Please enter your first name.");
            hasErrors = true;
        }

        // Validate Last Name
        if (!document.getElementById("client_last_name").value.trim()) {
            alert("Please enter your last name.");
            hasErrors = true;
        }

        // Validate Case Type Selection
        if (document.getElementById("type").value === "Please Select") {
            alert("Please select a case type.");
            hasErrors = true;
        }

        // Validate Detailed Description
        if (!document.getElementById("description").value.trim()) {
            alert("Please type a detailed description of your issue.");
            hasErrors = true;
        }

        // Validate Phone Number (Simple validation for demonstration purposes)
        const phoneNumber = document.getElementById("phone_number").value.trim();
        if (!phoneNumber || !/^\d{10}$/.test(phoneNumber)) {
            alert("Please enter a valid 10-digit phone number without any dashes or spaces.");
            hasErrors = true;
        }

        // Validate Email Address (Simple validation for demonstration purposes)
        const emailAddress = document.getElementById("email_address").value.trim();
        if (!emailAddress || !/\S+@\S+\.\S+/.test(emailAddress)) {
            alert("Please enter a valid email address.");
            hasErrors = true;
        }

        if (hasErrors) {
            event.preventDefault(); // Prevent form from submitting
            return false;
        }

        // If no errors, form will be submitted
    });
});


