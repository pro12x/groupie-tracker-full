document.addEventListener("DOMContentLoaded", function() {
    document.querySelector(".container").style.display = "none";
    setTimeout(function() {
        // Hide the preloader
        document.querySelector(".preloader-container").style.display = "none";
        // Display page content
        document.querySelector(".container").style.display = "block";
    }, 1000);
})

// Tabs
function openCity(evt, cityName) {
    var i, x, tablinks;
    x = document.getElementsByClassName("city");
    for (i = 0; i < x.length; i++) {
        x[i].style.display = "none";
    }
    tablinks = document.getElementsByClassName("tablink");
    for (i = 0; i < x.length; i++) {
        tablinks[i].className = tablinks[i].className.replace(" w3-border-blue", "");
    }
    document.getElementById(cityName).style.display = "block";
    evt.currentTarget.firstElementChild.className += " w3-border-blue";
}

// Sélectionner l'onglet par défaut au chargement de la page
window.onload = function() {
    var defaultTab = document.querySelector('.default-tab');
    if (defaultTab) {
        defaultTab.click();
    }
};

// Search function
function Search() {
    var search = document.querySelector('#search').value
    var result = document.querySelector('#result')
    const api = `/search?q=${encodeURIComponent(search)}`

    if(search === "") {
        result.style.display = "none";
        return
    }

    // fetch("/search?q=" + encodeURIComponent(search))
    fetch(api)
        .then((response) => {
            if (response.status === 204) {
                result.innerHTML = "No content found for " + search
                result.style.display = "block"
                // return
            }
            return response.json()
        }).then((data) => {
            if(!data) {
                return
            }

            // var result = document.querySelector('#result');
            result.innerHTML = ""
            result.style.display = "block"
            result.scrollTop = 0

            /*document.addEventListener("click", function (event) {
                if (!result.contains(event.target)) {
                    result.style.display = "none"
                    document.querySelector("#search").value = ""
                }
            })*/
            document.addEventListener("click", function (event) {
                if (!result.contains(event.target)) {
                    result.style.display = "none"
                    document.querySelector("#search").value = ""
                }
            })

            function Category(title) {
                var div = document.createElement("div")
                div.textContent = title
                div.style.fontWeight = "bold"
                div.style.textAlign = "center"
                div.style.color = "red"
                div.style.backgroundColor = "green"
                return div
            }

            function Result(suggestion) {
                var div = document.createElement("div");
                div.textContent = suggestion.value;
                div.onclick = function () {
                    window.location.href = "/artist/" + suggestion.id;
                };
                return div;
            }

            function Suggestion(category, resutls) {
                if(resutls !== null) {
                    resutls.sort(function (a, b) {
                        return a.value.localeCompare(b.value)
                    })

                    result.appendChild(
                        Category(category)
                    )

                    for(let i = 0; i < resutls.length; i++) {
                        result.appendChild(
                            Result(resutls[i])
                        )
                    }
                }
            }

            // Displaying suggestions
            Suggestion("Artist/Band", data.artists)
            Suggestion("Members", data.members)
            Suggestion("Locations", data.locations)
            Suggestion("First Albums", data.firstAlbums)
            Suggestion("Creation Dates", data.creationDates)
        }
    )
}