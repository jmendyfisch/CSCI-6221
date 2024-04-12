
// Mock data fetching function
async function fetchCases(assigned) {
    const lId = GLOBALS.lawyerId;
    let assignedInt = 0;
    let queryParams = { lawyer_id: 1 }; 

    if (assigned) {
        assignedInt = 1;
        queryParams = { lawyer_id: lId };
    }

    const url = '/cases?' + new URLSearchParams(queryParams);

    try {
        const response = await fetch(url);

        if (!response.ok) {
            console.log("error receiving data");
            return [];
        }

        const data = await response.json();

        if (!data) {
            return [];
        }

        const formattedData = data.map(obj => ({
            lastName: obj.client_last_name,
            firstName: obj.client_first_name,
            caseType: obj.type,
            dateInitiated: obj.created_at.split('T')[0],
            caseId: obj.id,
            assigned: assignedInt
        }));

        return formattedData;
    } catch (error) {
        console.error('There was a problem with the fetch operation:', error);
        return [];
    }
}

function createLink(caseId, text) {
    const link = document.createElement('a');
    link.href = `/case-details/${caseId}`;
    link.classList.add('table-link');
    link.textContent = text;
    return link.outerHTML;
}

function generateTableRow(caseInfo) {
    const row = document.createElement('tr');
    row.innerHTML = `
        <td>${createLink(caseInfo.caseId, caseInfo.lastName)}</td>
        <td>${createLink(caseInfo.caseId, caseInfo.firstName)}</td>
        <td>${createLink(caseInfo.caseId, caseInfo.caseType)}</td>
        <td>${createLink(caseInfo.caseId, caseInfo.dateInitiated)}</td>
        <td><button class="assign-btn" data-case-id="${caseInfo.caseId}">${caseInfo.assigned ? 'Detach' : 'Assign'}</button></td>
    `;
    return row;
}

async function populateTable(containerId, cases) {
    // Use containerId to select the correct table body within that container
    const tableBody = document.querySelector(`#${containerId} table tbody`);
    if (!tableBody) {
        console.error(`Table body not found for container: ${containerId}`);
        return;
    }

    cases.forEach(caseInfo => {
        tableBody.appendChild(generateTableRow(caseInfo));
    });

    // Add event listeners for assign/detach buttons
    document.querySelectorAll(`#${containerId} .assign-btn`).forEach(button => {
        button.addEventListener('click', function() {
            const caseId = this.getAttribute('data-case-id');
            if (this.textContent === 'Assign') {
                // Placeholder: Implement case assignment logic
                console.log(`Assigning case ID: ${caseId}`);
            } else {
                // Placeholder: Implement case detachment logic
                console.log(`Detaching case ID: ${caseId}`);
            }
        });
    });
}


// Sorting functionality
function sortTable(tableId, column) {
    const getCellValue = (tr, idx) => tr.children[idx].innerText || tr.children[idx].textContent;
    const comparer = (idx, asc) => (a, b) => ((v1, v2) => 
        v1 !== '' && v2 !== '' && !isNaN(v1) && !isNaN(v2) ? v1 - v2 : v1.toString().localeCompare(v2)
    )(getCellValue(asc ? a : b, idx), getCellValue(asc ? b : a, idx));

    const table = document.querySelector(`#${tableId} table`);
    const tbody = table.querySelector('tbody');
    Array.from(tbody.querySelectorAll('tr'))
          .sort(comparer(Array.from(tbody.parentNode.querySelectorAll('th')).indexOf(column), this.asc = !this.asc))
          .forEach(tr => tbody.appendChild(tr));
}

document.addEventListener('DOMContentLoaded', async () => {

    const assignedCases = await fetchCases(true); // Fetch assigned cases
    const unassignedCases = await fetchCases(false); // Fetch unassigned cases (this might be a different fetch in a real application)
    await populateTable('assignedCases', assignedCases); // Use the ID of the div, not the tbody
    await populateTable('unassignedCases', unassignedCases); // Use the ID of the div, not the tbody

    // Attach sorting functionality to headers
    document.querySelectorAll('#assignedCases th').forEach(header => {
        header.addEventListener('click', () => sortTable('assignedCases', header));
    });
    document.querySelectorAll('#unassignedCases th').forEach(header => {
        header.addEventListener('click', () => sortTable('unassignedCases', header));
    });
});

document.addEventListener('DOMContentLoaded', () => {
    // Query all table rows
    const rows = document.querySelectorAll('table tr');

    // Add hover event listeners to each row
    rows.forEach(row => {
        const links = row.querySelectorAll('a.table-link');
        links.forEach(link => {
            // Mouse enter event
            link.addEventListener('mouseenter', () => {
                row.classList.add('hovered-row');
            });
            // Mouse leave event
            link.addEventListener('mouseleave', () => {
                row.classList.remove('hovered-row');
            });
        });
    });
});
