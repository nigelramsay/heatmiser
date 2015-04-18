package heatmiser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeatMiser(t *testing.T) {
	var h = HeatMiser{}
	h.body = []byte(ExampleResponse)

	assert.Equal(t, 19.8, h.Current(), "they should be equal")

	assert.Equal(t, 15.0, h.Target(), "they should be equal")

	assert.Equal(t, false, h.Enabled(), "they should be equal")

}

const ExampleResponse = `
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">
<html>
<head>
<title>Heatmiser Wifi Thermostat</title>
<link href="/mywifi.css" rel="stylesheet" type="text/css" />
	<meta http-equiv="Content-Type" content="text/html; charset=us-ascii">
	<meta http-equiv="Content-Language" content="en-gb">
	<meta http-equiv="Cache-Control" content="no-cache, must-revalidate">
</head>

<script language="javascript">
function awayClick()
{	aw = document.FrmChg4;
	if(aw.actH.value=='1')
	{
		aw.actH.value=2;
	}
	else
	{
		aw.actH.value=1;
	}
	aw.submit();
}
function summerClick()
{	sm = document.fostfm;
	if(sm.fost.value=='1')
	{
		sm.fost.value=2;
	}
	else
	{
		sm.fost.value=1;
	}
	sm.submit();
}
function myrefsh()
{
	location.assign(location);
}
function myonload()
{}
</script>
<body bgcolor="#B3120C" onload="myonload()" topmargin=0 leftmargin=0 rightmargin=0 style="overflow-x:hidden;overflow-y:auto">
		<form name="logfm">
		<input type=hidden name="lgst" value="1">
		<input type=hidden name="disw" value="none">
		<input type=hidden name="modl" value="2">
		<input type=hidden name="senor" value="0">
		<input type="hidden" name="hldy" value="2">
		</form>
<script language="javascript">
		if(document.logfm.lgst.value != "1")
			top.location.href="index.htm";
</script>
<form name="FrmChg4" method="post">
<input type="hidden" name="actH" value="2">
</form>
<form name="fostfm" method="post">
<input type="hidden" name="fost" value="2">
<form>
		<div id="lb7" ></div>
		<form name='dispFrm'>
		<table bgcolor="#B3120C" style='text-align:center;' border="0" width=100% height=100%  cellspacing=0 cellpadding=0 >
		<script language="javascript">
			var tmp, sntm, hltm;
			tmp = Number(document.logfm.modl.value);
			sntm = Number(document.logfm.senor.value);
			hltm = Number(document.logfm.hldy.value);
			if(tmp > 1)
				document.write("<tr><td align=center height=30 colspan=4 ><a href='/statSetup.htm' target='midfm' onfocus='this.blur()' ><b><font color='white' face='arial'>10:21&nbsp;&nbsp;19/04/2015</font></b></a></td></tr>");
			else
				document.write("<tr><td align=center height=30 colspan=4 ><br></td></tr>");
			document.write("<tr><td height=30 bgcolor='black' colspan=4></td></tr><td class='p5' colspan=4 ><b>Live View&nbsp;</b><input type='button' value='Refresh' onclick='myrefsh()'></td></tr><tr><td colspan=4 ><br></td></tr>");
			if(tmp < 5)
			{
				if(sntm > 1)
					document.write("<tr><td class='p5'colspan=4><b>Floor&nbsp;:&nbsp;</b><font size='5'>&nbsp;<font face='Arial'>&deg;</font>C</font></td></tr>");
				if(sntm != 2)
					document.write("<tr><td class='p5'colspan=4><b>Actual&nbsp;:&nbsp;</b><font size='5'>19.8&nbsp;<font face='Arial'>&deg;</font>C</font></td></tr><tr><td class='p5'colspan=4><b>Set&nbsp;:&nbsp;</b><font size='4'>15&nbsp;<font face='Arial'>&deg;</font>C</font></td></tr><tr><td colspan=4 ><br></td></tr><tr><td class='p5'colspan=4><b>Heat Status:</b><font size='4'>OFF&nbsp;<font face='Arial'></td></tr>");
			}
			if(tmp > 3)
		 	{
		 		if(tmp < 5 )
		 			document.write("<tr><td class='p5'colspan=4><b>Hot Water:</b><font size='4'>&nbsp;<font face='Arial'></td></tr>");
				else
					document.write("<tr><td class='p5'colspan=4><b>Timer :</b><font size='4'>&nbsp;<font face='Arial'></td></tr>");

				if(tmp == 6)
					document.write("<tr><td colspan=4 ><br></td></tr><tr><td class='p5'colspan=4><b>Timer left&nbsp;: 42 &nbsp;h:&nbsp;40&nbsp;min</b></td></tr>");
	 		}

			if((tmp > 1)&&(tmp < 5))
				document.write("<tr><td colspan=4 ><br></td></tr><tr><td class='p5'colspan=4><b>Hold&nbsp;for:&nbsp;0&nbsp;h:&nbsp;00&nbsp;min</b></td></tr>");
			document.write("<tr><td colspan=4 ><br></td></tr><tr><td class='p5' colspan=4><b>Occupancy</b></td></tr>");
			document.write("<tr><td class='p5'colspan=1 >&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</td><td class='p5' style='text-decoration:underline;'colspan=1><b>Home</b></td><td class='p5' colspan=1><a onclick=awayClick() style='color:white;' href='#'><b>Away</b></a></td><td class='p5'colspan=1 >&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</td></tr>");
			if(tmp == 4)
			{
				if(hltm!=1)
				{
					document.write("<tr><tr><td colspan=4 ><br></td></tr><tr><td class='p5' colspan=4><b>Summer</b></td></tr>");
					document.write("<tr><td class='p5'colspan=1 >&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</td><td class='p5' </td><td class='p5'  </td><td class='p5'colspan=1 >&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</td></tr>");
				}
			}
		</script>
		</table>
		</form>
	</body>
</html>`
