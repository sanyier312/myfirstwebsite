
let c=1;
let d=1;
let e=1;
function f1()
{ c++;
  if(c%2)
  {
    document.getElementById('hid1').readOnly=false;
  document.getElementById('wr1').innerText='保存';
  }
  else{document.getElementById('hid1').readOnly=true;
  document.getElementById('wr1').innerText='编辑';
  }

}
function f2()
{ c++;
  if(c%2)
  {
    document.getElementById('hid2').readOnly=false;
  document.getElementById('wr2').innerText='保存';
  }
  else{document.getElementById('hid2').readOnly=true;
  document.getElementById('wr2').innerText='编辑';
  }

}
function f3()
{ c++;
  if(c%2)
  {
    document.getElementById('hid3').readOnly=false;
  document.getElementById('wr3').innerText='保存';
  }
  else{document.getElementById('hid3').readOnly=true;
  document.getElementById('wr3').innerText='编辑';
  }

}
 function g1()
{
  d++;
  if(d%2)
  {
    document.getElementById('l1').innerText='已喜欢';
    document.getElementById('l1').style.background='red';
  }
  else
  {
    document.getElementById('l1').innerText='喜欢';
    document.getElementById('l1').style.background='white';

  }
}
function g2()
{
  d++;
  if(d%2)
  {
    document.getElementById('l2').innerText='已喜欢';
    document.getElementById('l2').style.background='red';
  }
  else
  {
    document.getElementById('l2').innerText='喜欢';
    document.getElementById('l2').style.background='white';

  }
}
function g3()
{
  d++;
  if(d%2)
  {
    document.getElementById('l3').innerText='已喜欢';
    document.getElementById('l3').style.background='red';
  }
  else
  {
    document.getElementById('l3').innerText='喜欢';
    document.getElementById('l3').style.background='white';

  }
}
function h1(){
  
     document.getElementById('td1').style.display='block';
     document.getElementById('pl1').style.width="900px";

 
}
function h2(){
  
  document.getElementById('td2').style.display='block';
  document.getElementById('pl2').style.width="900px";


}
function h3(){
  
  document.getElementById('td3').style.display='block';
  document.getElementById('pl3').style.width="900px";


}
function i()
{

  var name = document.getElementById("uid").innerText
  document.getElementById("name").value = name
  var ct = document.forms[2];
  ct.submit();

}
function j()
{
  location.href="https://wx.qq.com/"
}
function k()
{
  var aname = document.getElementById("uid").innerText
  document.getElementById("aname").value = aname
  var f =  document.forms[1];
  f.submit();
}
