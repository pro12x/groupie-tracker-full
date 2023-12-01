document.addEventListener("DOMContentLoaded", function() {
    document.querySelector(".container").style.display = "none";
    setTimeout(function() {
        // Hide the preloader
        document.querySelector(".preloader-container").style.display = "none";
        // Display page content
        document.querySelector(".container").style.display = "block";
    }, 1000);
})

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