var navItems = document.querySelectorAll("nav .nav-item");

navItems.forEach(navItem => {
    if(navItem) {
        navItem.onclick = () => {
            navItems.forEach(navItem => {
                navItem.classList.remove("active");
            })
            navItem.classList.add("active");
        }
    }
})

/* dropdown functionality */

document.addEventListener("click", e => {
    const isDropdownBtn = e.target.matches("[data-dropdown-btn]");
    if(!isDropdownBtn && e.target.closest("[data-dropdown]") != null){
        return;
    }

    let currentDropdown
    if(isDropdownBtn) {
        currentDropdown = e.target.closest("[data-dropdown]");
        currentDropdown.classList.toggle("menu-open");
    }

    document.querySelectorAll("[data-dropdown].menu-open").forEach(dropdown => {
        var clsDropdownBtn = dropdown.querySelector("[data-dropdown-cls]");
        if(clsDropdownBtn) {
            clsDropdownBtn.onclick = () => {
                dropdown.classList.remove("menu-open");
            }
        }
        if(dropdown === currentDropdown){
            return;
        }
        dropdown.classList.remove("menu-open");
        
    })

    
})

/* tabs functionality */

var tabs = document.querySelector(".tab-bar");
var tabButton = document.querySelectorAll(".tab-button");
var contents = document.querySelectorAll(".tab-content-wrap > .tab-content");

if (tabs) {
    tabs.onclick = e => {
        var id = e.target.dataset.id;
        console.log(id);
        if (id) {
            tabButton.forEach(btn => {
                btn.classList.remove("active-tab");
            });
            e.target.classList.add("active-tab");

            contents.forEach(content => {
                content.classList.remove("active-content");
            });
            var element = document.getElementById(id);
            element.classList.add("active-content");
        }
    }
}


var innerTabWrapper = document.querySelectorAll(".tab-content-wrap > .tab-content")
innerTabWrapper.forEach(innerTabWrap => {
    var innerTabs = innerTabWrap.querySelector(".inner-tab-bar");
    var innerTabButton = innerTabWrap.querySelectorAll(".inner-tab-button");
    var innerContents = innerTabWrap.querySelectorAll(".inner-tab-content-wrap .tab-content");

    if (innerTabs) {
        innerTabs.onclick = e => {
            var id = e.target.dataset.id;
            if (id) {
                innerTabButton.forEach(btn => {
                    btn.classList.remove("active-tab");
                });
                e.target.classList.add("active-tab");
                

                innerContents.forEach(content => {
                    content.classList.remove("active-content");
                });
                var element = document.getElementById(id);
                element.classList.add("active-content");
            }
        }
    }
    
})

/* message */

var message = document.querySelector("[data-message]");


if(message) {
    var clsMessage = message.querySelector("[data-close-message]");
    clsMessage.onclick = () => {
        message.classList.add("hide");
    }
}