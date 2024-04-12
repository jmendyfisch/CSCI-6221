// Mock data fetching functions
async function fetchClientInfo(caseId) {
    // Placeholder: Fetch client info from your backend here
    return {
        name: 'John Doe',
        phoneNumber: '123-456-7890',
        emailAddress: 'johndoe@example.com',
        streetAddress: '123 Main St, Anytown, AN 12345',
        caseDescription: 'Description of the case...'
    };
}

async function fetchMeetings(caseId) {
    // Placeholder: Fetch meetings from your backend here
    return [
        { meetingTime: '2024-04-12T10:00:00Z', lawyerNotes: 'Initial consultation', meetingId: 1 },
        { meetingTime: '2024-04-19T14:00:00Z', lawyerNotes: 'Follow-up meeting', meetingId: 2 }
    ];
}

function createLink(meetingId, text) {
    const link = document.createElement('a');
    link.href = `/meeting_details/${meetingId}`;
    link.classList.add('table-link');
    link.textContent = text;
    return link.outerHTML;
}

function populateClientInfo(clientInfo) {
    const tableBody = document.querySelector('#clientInfo table tbody');
    tableBody.innerHTML = `
        <tr><td>Name</td><td>${clientInfo.name}</td></tr>
        <tr><td>Phone Number</td><td>${clientInfo.phoneNumber}</td></tr>
        <tr><td>Email Address</td><td>${clientInfo.emailAddress}</td></tr>
        <tr><td>Street Address</td><td>${clientInfo.streetAddress}</td></tr>
        <tr><td>Case Description</td><td>${clientInfo.caseDescription}</td></tr>
    `;
}

function generateTableRow(meetingInfo) {
    const row = document.createElement('tr');
    row.innerHTML = `
        <td>${createLink(meetingInfo.meetingId, meetingInfo.meetingTime)}</td>
        <td>${meetingInfo.lawyerNotes}</td>
    `;
    return row;
}

async function populateMeetings(caseId) {
    const meetings = await fetchMeetings(caseId);
    const tableBody = document.querySelector('#meetingsBody');
    meetings.forEach(meetingInfo => {
        tableBody.appendChild(generateTableRow(meetingInfo));
    });
}

document.addEventListener('DOMContentLoaded', async () => {
    const urlParams = new URLSearchParams(window.location.search);
    const caseId = urlParams.get('caseId');
    
    const clientInfo = await fetchClientInfo(caseId);
    populateClientInfo(clientInfo);
    
    await populateMeetings(caseId);
});
