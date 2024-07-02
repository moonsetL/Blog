document.write("<span id=time></span>")
setInterval("time.innerText=new Date().toLocaleString()",1000)