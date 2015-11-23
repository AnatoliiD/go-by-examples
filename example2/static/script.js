$(function(){
  $(document).on('click', '#make_request', function(e){
    console.log();
    $.ajax({
      url: $(this).attr('href'),
      method: 'GET',
      success: function(d,e,v){
        console.log(d);
        console.log(e);
        console.log(v);
        $('#result').text(v.responseText);
      }
    })
    return false;
  });
});
