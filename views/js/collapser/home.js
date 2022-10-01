
let collapser = document.getElementById("collapser")
let collapsericon=document.getElementById("collapsericon")
let counter=0
let leftbar = document.getElementById("leftbar")
collapser.addEventListener('click',function(e){
e.preventDefault()
if (counter == 0) {

collapsericon.classList.add("fa-close")
collapsericon.classList.remove('fa-bars')
leftbar.style="left:0px"

counter=1
}else{
    collapsericon.classList.remove("fa-close")
    collapsericon.classList.add('fa-bars')
    leftbar.style="left:-210px"
    counter=0
}


})
document.documentElement.requestFullscreen().then(()=>console.log("full screen"));