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