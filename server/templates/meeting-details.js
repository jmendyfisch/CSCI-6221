// Mock data fetching functions
async function fetchMeetingInfo(caseId) {
    // Placeholder: Fetch meeting info from your backend here
    return {
        meetingTime: '2024-04-12T10:00:00',
        clientFirstName: 'John',
        clientLastName: 'Doe',
        phoneNumber: '123-456-7890',
        email: 'johndoe@example.com',
        lawyerNotes: 'Initial meeting for case review.'
    };
}

async function fetchGptResponses(caseId) {
    // Placeholder: Fetch GPT responses from your backend here
    return [
        { id: 1, question: 'What is the case about? What actions have been taken?', summary: 'Case regarding property dispute.', points: 'Ask the client about his lease' },
        { id: 2, question: 'How are you feeling today? How much is your rent?' , summary: 'Initial documentation filed.', points: 'Ask what court it is' }
    ];
}

function populateMeetingInfo(meetingInfo) {
    const tableBody = document.querySelector('#meetingDetailsBody');
    tableBody.innerHTML = `
        <tr>
            <td>${meetingInfo.meetingTime}</td>
            <td>${meetingInfo.clientFirstName}</td>
            <td>${meetingInfo.clientLastName}</td>
            <td>${meetingInfo.phoneNumber}</td>
            <td>${meetingInfo.email}</td>
            <td>${meetingInfo.lawyerNotes}</td>
        </tr>
    `;
}

function generateGptResponseRow(response, index) {
    const row = document.createElement('tr');
    row.innerHTML = `
        <td>${index}</td>
        <td>${response.question}</td>
        <td>${response.summary}</td>
        <td>${response.points}</td>
    `;
    return row;
}

async function populateGptResponses(caseId) {
    const responses = await fetchGptResponses(caseId);
    const tableBody = document.querySelector('#gptResponsesBody');
    responses.forEach((response, index) => {
        tableBody.appendChild(generateGptResponseRow(response, index + 1));
    });
}

document.addEventListener('DOMContentLoaded', async () => {
    const urlParams = new URLSearchParams(window.location.search);
    const caseId = urlParams.get('caseId');
    
    const meetingInfo = await fetchMeetingInfo(caseId);
    populateMeetingInfo(meetingInfo);
    
    await populateGptResponses(caseId);
});
