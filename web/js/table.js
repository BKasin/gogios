function replaceText(title, output) {
  document.getElementById('CheckName').textContent = title;

  document.getElementById('CheckOutput').innerHTML = "<pre>" + output + "</pre>";
}

$(document).ready(function () {
  $.getJSON("js/current.json",
    function (json) {
      var tr;
      // Append each row to html table
      for (var i = 0; i < json.length; i++) {
        tr = $('<tr/>');
        tr.append("<td><a href='#' onClick='replaceText(\"" + json[i].title + "\"," + json[i].output + ");'>" + json[i].title + "</a></td>");
        if (json[i].good == true) {
          tr.append("<td><font color='green'>Success</font></td>");
        }
        else if (json[i].good == false) {
          tr.append("<td><font color='red'>Fail</font></td>");
        }
        tr.append("<td>" + json[i].asof + "</td>");
        $('table').append(tr);
      }
    });
});
