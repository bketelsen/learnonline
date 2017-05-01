require('expose-loader?$!expose-loader?jQuery!jquery');

$(() => {
    activateSideNav();
});

function activateSideNav() {
    let loc = window.location;
    let path = loc.pathname;
    let base = '/' + path.split('/').slice(1,3).join('/');
    $("ul.nav li").removeClass("active");
    $("#livideo").addClass("active");
    $(`ul.nav a[href='${base}']`).closest("li").addClass("active");
    $(`ul.nav a[href='${base}']`).closest("div").addClass("collapse in");
}


