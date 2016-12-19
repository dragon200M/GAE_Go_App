
var d =[]
$(document).ready(function() {

    $(function(){

    $.ajax({
        url: "/serve/time",
        type: "get",
        success: function (result) {
           var  c = result.split(' ');

           for(var i=0;i< c.length;i++){
            d.push(c[i])

           }

            d = c
        }
    });

});
console.log(d)


var monthNames = [ "Styczeń", "Luty", "Marzec", "Kwiecień", "Maj", "Czerwiec", "Lipiec", "Sierpień", "Wrzesień", "Październik", "Listopad", "Grudzień" ]; 
var dayNames= ["Niedziela","Poniedziałek","Wtorek","Środa","Czwartek","Piątek","Sobota"]


var newDate = new Date();

newDate.setDate(newDate.getDate());

  
$('#Date').html(dayNames[newDate.getDay()] + " " + newDate.getDate() + ' ' + monthNames[newDate.getMonth()] + ' ' + newDate.getFullYear());

setInterval( function() {
	
	var seconds = new Date().getSeconds();
	
	$("#sec").html(( seconds < 10 ? "0" : "" ) + seconds);
	},1000);
	
setInterval( function() {
	
	var minutes = new Date().getMinutes();
	
	$("#min").html(( minutes < 10 ? "0" : "" ) + minutes);
    },1000);
	
setInterval( function() {
	
	var hours = new Date().getHours();

	$("#hours").html(( hours < 10 ? "0" : "" ) + hours);
    }, 1000);	
});