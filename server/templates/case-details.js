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
        { meetingTime: '2024-04-12T10:00:00Z', lawyerNotes: 'Initial consultation', meetingId: 1},
        { meetingTime: '2024-04-19T14:00:00Z', lawyerNotes: 'Follow-up meeting', meetingId: 2 }
    ];
}

function createLink(meetingId, text) {
    const link = document.createElement('a');
    
    link.href = `/meeting-details/${GLOBALS.caseId}/${meetingId}`;
    link.classList.add('table-link');
    link.textContent = text;
    return link.outerHTML;
}

function populateClientInfo(clientInfo) {
    const tableBody = document.querySelector('#clientInfo table tbody');
    tableBody.innerHTML = `
        <tr><td>${clientInfo.name}</td><td>${clientInfo.phoneNumber}</td><td>${clientInfo.emailAddress}</td><td>${clientInfo.streetAddress}</td></tr>
        <tr><td colspan="4"><b>Client case description:</b> ${clientInfo.caseDescription}</td></tr>
        <tr><td colspan="4"><b>AI-generated case summary:</b> [Your AI Summary Here]</td></tr>
    `;
}


function generateTableRow(meetingInfo) {
    // Convert meetingTime to desired format here
    const meetingDate = new Date(meetingInfo.meetingTime);
    const formattedDate = meetingDate.toLocaleDateString('en-US', {
        month: 'short', day: '2-digit', year: 'numeric', hour: '2-digit', minute: '2-digit'
    });

    const row = document.createElement('tr');
    row.innerHTML = `
        <td class="meeting-time-column">${createLink(meetingInfo.meetingId, formattedDate)}</td>
        <td class="lawyer-notes-column">${meetingInfo.lawyerNotes}</td>
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



