
document.getElementById("startIntakeBtn").addEventListener("click", function() {
    document.getElementById("intakeForm").style.display = "block";
    this.style.display = "none";
});

document.getElementById("intakeForm").addEventListener("submit", function(event) {
    event.preventDefault();

    const formData = {
        client_name: document.getElementById("client_name").value,
        type: document.getElementById("type").value,
        description: document.getElementById("description").value,
        contact: document.getElementById("contact").value,
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
