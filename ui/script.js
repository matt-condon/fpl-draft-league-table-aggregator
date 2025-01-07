// Fetch the aggregated league table data from the JSON file
fetch(filePath)
    .then(response => response.json())
    .then(data => {
        const tableBody = document.getElementById('leagueTable').getElementsByTagName('tbody')[0];

        data.Standings.forEach(entry => {
            const row = document.createElement('tr');

            // Add table cells
            row.innerHTML = `
                <td>${entry.Rank}</td>
                <td>
                    <a href="${entry.TeamUrl}">${entry.EntryName}</a>
                    <div>${entry.PlayerName}</div>
                </td>
                <td>${entry.EventTotal}</td>
                <td>${entry.StageTwoTotal}</td>
                <td>${entry.Total}</td>
            `;

            tableBody.appendChild(row);
        });

        const eventElement = document.getElementById('event');
        eventElement.textContent = `[${data.Event}]`;
    })
    .catch(error => console.error('Error fetching league table data:', error));

function sortTable(columnIndex) {
    const table = document.getElementById("leagueTable");
    const rows = Array.from(table.tBodies[0].rows);
    const isNumeric = !isNaN(rows[0].cells[columnIndex].innerText.trim());
    const currentDirection = table.getAttribute("data-sort-direction") || "asc";
    const newDirection = currentDirection === "asc" ? "desc" : "asc";

    // Sort rows
    rows.sort((rowA, rowB) => {
        const cellA = rowA.cells[columnIndex].innerText.trim();
        const cellB = rowB.cells[columnIndex].innerText.trim();

        if (isNumeric) {
            return newDirection === "asc" 
                ? parseFloat(cellA) - parseFloat(cellB) 
                : parseFloat(cellB) - parseFloat(cellA);
        } else {
            return newDirection === "asc" 
                ? cellA.localeCompare(cellB) 
                : cellB.localeCompare(cellA);
        }
    });

    // Reattach sorted rows
    const tbody = table.tBodies[0];
    tbody.innerHTML = "";
    rows.forEach(row => tbody.appendChild(row));

    // Update sort direction
    table.setAttribute("data-sort-direction", newDirection);
    table.setAttribute("data-sorted-column", columnIndex);

    // Reset all headers and set indicator for the sorted column
    const headers = table.querySelectorAll("th");
    headers.forEach((header, idx) => {
        header.classList.remove("sorted-asc", "sorted-desc");
        if (idx === columnIndex) {
            header.classList.add(newDirection === "asc" ? "sorted-asc" : "sorted-desc");
        }
    });
}
