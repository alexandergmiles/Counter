$(document).ready(function(){
    increaseCount();
});

function increaseCount() {
    $.ajax({
        url: "IncreaseCount",
        success: function(result) {
            $(".count").text(result);
        }
    })
}
