// Mock data fetching functions

async function fetchClientInfo(caseId) {

    const cId = GLOBALS.caseId;
    let queryParams2 = {case_id: cId};

    const url2 = '/get-case-details?' + new URLSearchParams(queryParams2);

    try {
        const response2 = await fetch(url2);

        if (!response2.ok) {
            console.log("error receiving data");
            return {};
        }

        const caseData = await response2.json();

        if (!caseData) {
            return {};
        }
        //console.log("caseData: " + caseData);
        let formattedData = {
            clientName: caseData.client_first_name+" "+caseData.client_last_name,
            phoneNumber: caseData.phone_number,
            email: caseData.email_address,
            address: caseData.address_street+", "+caseData.address_city+", "+caseData.address_state+" "+caseData.address_zip,
            description: caseData.description,
            gptsummary: caseData.gpt_summary
        };


        return formattedData;
    } catch (error) {
        console.error('There was a problem with the fetch operation:', error);
        return [];
    }
}




async function fetchMeetings(caseId) {
    

    const cId = GLOBALS.caseId;
    let queryParams1 = { case_id: cId}; 

    const url1 = '/get-all-meetings?' + new URLSearchParams(queryParams1);

    try {
        const response1 = await fetch(url1);

        if (!response1.ok) {
            console.log("error receiving data");
            return [];
        }

        const meetData = await response1.json();

        if (!meetData) {
            return [];
        }

        let formattedData = meetData.map(obj => ({
            meetingId: obj.id,
            meetingTime: obj.created_at,
            lawyerNotes: obj.lawyer_notes.String
        }));

        return formattedData;
    } catch (error) {
        console.error('There was a problem with the fetch operation:', error);
        return [];
    }
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
    summaryText = "";
    if (clientInfo.gptsummary != ""){
        summaryText = clientInfo.gptsummary;
        //console.log("summaryText: " + summaryText);
    }else{
        summaryText = "Will be created upon completion of first meeting.";
        //console.log("summaryText: " + summaryText);
    }

    tableBody.innerHTML = `
        <tr><td>${clientInfo.clientName}</td><td>${clientInfo.phoneNumber}</td><td>${clientInfo.email}</td><td>${clientInfo.address}</td></tr>
        <tr><td colspan="4"><b>Client description:</b> ${clientInfo.description}</td></tr>
        <tr><td colspan="4"><b>AI-generated summary:</b> ${summaryText}</td></tr>
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


async function populateMeetings(caseId, sortedMeetings = null) {
    const meetings = sortedMeetings ? sortedMeetings : await fetchMeetings(caseId);
    const tableBody = document.querySelector('#meetingsBody');
    const meetingsSection = document.getElementById('meetingDetails'); // Get the entire meetings section

    // Clear existing rows before re-populating
    tableBody.innerHTML = '';

    if (meetings.length > 0) {
        meetings.forEach(meetingInfo => {
            tableBody.appendChild(generateTableRow(meetingInfo));
        });
        // Make sure the meetings section is visible if there are meetings
        meetingsSection.style.display = ''; // Use '' to reset to default display style
    } else {
        // If there are no meetings, hide the entire meetings section
        meetingsSection.style.display = 'none';
    }
}




document.addEventListener('DOMContentLoaded', async () => {

    const urlParams = new URLSearchParams(window.location.search);
    const caseId = urlParams.get('caseId');
    const clientInfo = await fetchClientInfo(caseId);
    populateClientInfo(clientInfo);

    let meetings = await fetchMeetings(caseId); // Fetch meetings initially

    populateMeetings(caseId, meetings); // Populate with initial unsorted data

    // Find the "Meeting Time" header and attach click event
    document.querySelector('.meeting-time-column').addEventListener('click', async () => {
        meetings = sortMeetingsByTime(meetings); // Sort meetings
        populateMeetings(caseId, meetings); // Re-populate table with sorted data
    });
});



let sortAscending = true; // Keep track of sort direction

function sortMeetingsByTime(meetings) {
    // Sort meetings by the 'meetingTime' property
    meetings.sort((a, b) => {
        const dateA = new Date(a.meetingTime);
        const dateB = new Date(b.meetingTime);
        return sortAscending ? dateA - dateB : dateB - dateA;
    });

    // Toggle sort direction for next sort
    sortAscending = !sortAscending;
    return meetings;
}
