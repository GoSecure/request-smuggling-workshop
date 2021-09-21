<?php
$title = "Contact";
$page = $_SERVER['PHP_SELF'];
if(isset($_SERVER['QUERY_STRING'])) {
    $page .= "?".$_SERVER['QUERY_STRING'];
}

include_once("_header.php");
?>

<h1>Contact</h1>

<form action="<?= $page ?>" method="post">

Name: <input type="text" name="name" value="Your Name"><br/>
<br/>

Content: <br/>
<textarea name="content" cols="80" rows="5">
Your Message
</textarea><br/>
<br/>

<input type="submit" value="Submit">
</form>

<?php
include_once("_footer.php");
?>