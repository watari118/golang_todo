function validateTodoUpsert() {
  var content = document.getElementById("content").value.trim();
  if (content === "") {
    document.getElementById("error-message").style.display = "block";
    return false; // フォーム送信を防ぐ
  }
  document.getElementById("error-message").style.display = "none";
  return true; // フォーム送信を許可
}

function validateSignin() {
  var email = document.getElementById("email").value.trim();
  var password = document.getElementById("password").value.trim();
  var errorMessage = "";

  if (email === "" && password === "") {
    errorMessage = "Emailとパスワードを入力してください。";
  } else if (email === "") {
    errorMessage = "Emailを入力してください。";
  } else if (password === "") {
    errorMessage = "パスワードを入力してください。";
  }

  if (errorMessage !== "") {
    document.getElementById("error-message").innerText = errorMessage;
    document.getElementById("error-message").style.display = "block";
    return false; // フォーム送信を防ぐ
  }

  document.getElementById("error-message").style.display = "none";
  return true; // フォーム送信を許可
}

function confirmDelete() {
  return confirm("本当に削除しますか？");
}

function confirmLogout() {
  return confirm("ログアウトしますか？");
}
