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
var tabButton = document.querySelectorAll(".tab-bar > .tab-button");
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

var modalTabs = document.querySelector(".modal-tab-bar");
var modalTabButton = document.querySelectorAll(".modal-tab-bar .tab-button");
var modalTabContents = document.querySelectorAll(".modal-tab-content-wrap > .modal-tab-content");

if (modalTabs) {
    modalTabs.onclick = e => {
        var id = e.target.dataset.id;
        console.log(id);
        if (id) {
            modalTabButton.forEach(btn => {
                btn.classList.remove("active-tab");
            });
            e.target.classList.add("active-tab");

            modalTabContents.forEach(content => {
                content.classList.remove("active-content");
            });
            var element = document.getElementById(id);
            element.classList.add("active-content");

        }
    }
}

/* modal */

var modalButtons = document.querySelectorAll("[data-modal]");
// var modalContents = document.querySelectorAll(".modal-wrap");

modalButtons.forEach(modalButton => {
    if(modalButton) {
        modalButton.onclick = e => {
            var modalId = e.target.dataset.modal;
            console.log(modalId);

            var modalContent = document.getElementById(modalId);
            modalContent.classList.add("active-modal");
            document.querySelector("body").classList.add("modal-open");

            var closeModal = modalContent.querySelector("[data-modal-closer]");
            if(closeModal) {
                closeModal.onclick = () => {
                    modalContent.classList.remove("active-modal");
                    document.querySelector("body").classList.remove("modal-open");
                }
            }
        }
    }
})

/* collapsible */

var coll = document.getElementsByClassName("collapsible");
var i;

for (i = 0; i < coll.length; i++) {
    coll[i].addEventListener("click", function() {
        this.classList.toggle("active");
        var content = this.nextElementSibling;
        
        if (content.style.maxHeight){
            content.style.maxHeight = null;
        
        } else {
            content.style.maxHeight = content.scrollHeight + "px";
        
        } 
    });
};

/* message */

var message = document.querySelector("[data-message]");


if(message) {
    var clsMessage = message.querySelector("[data-close-message]");
    clsMessage.onclick = () => {
        message.classList.add("hide");
    }
}

/* copy clipboard */

document.addEventListener("click", e => {
    var isCopyBtn = e.target.matches("[data-copy-button]");
    if(!isCopyBtn && e.target.closest("[data-copy-clipboard]") != null){
        return;
    }

    let currentClipboard
    if(isCopyBtn) {
        currentClipboard = e.target.closest("[data-copy-clipboard]");
        var copyContent = currentClipboard.querySelector(".copy-content");

        // if (document.selection){
        //     var div = currentClipboard.createTextRange();

        //     div.moveToElementText(copyContent);
        //     copyContent.select();
        // } else {
        //     var div = document.createRange();

        //     div.setStartBefore(copyContent);
        //     div.setEndAfter(copyContent);

        //     window.getSelection().addRange(div);
        // }
        copyContent.select();    
        document.execCommand("Copy");
    }

    document.querySelectorAll("[data-copy-clipboard]").forEach(clipboard => {
        if(clipboard === currentClipboard){
            return;
        }
        clipboard.classList.remove("menu-open");
    })

    
});


/* custom range */
const rangeInputs = document.querySelectorAll('input[type="range"]');

function handleInputChange(e) {
    let target = e.target
    if (e.target.type !== 'range') {
        target = document.getElementById('range')
        console.log(target)
    } 
    const min = target.min
    const max = target.max
    const val = target.value

    
    target.style.backgroundSize = (val - min) * 100 / (max - min) + '% 100%'
}

rangeInputs.forEach(input => {
    input.addEventListener('input', handleInputChange)
});