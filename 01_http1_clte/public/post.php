<?php

if(isset($_POST['content'])) {
    error_log("Data:".$_POST['content']);
	echo $_POST['content'];
}
else {
	error_log("No content received");
}





?>