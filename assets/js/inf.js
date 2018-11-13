$(document).ready(function() {
    $.ajax({
        url:"/jsTest"
    }).then(function(data) {
        $('.username').append(data.username);
        $('.password').append(data.password);
    });
})