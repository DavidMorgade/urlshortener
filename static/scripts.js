document.getElementById('url-form').addEventListener('submit', async function(event) {
  const domain = window.location.origin;
  event.preventDefault();
  const urlInput = document.getElementById('url-input').value;
  const response = await fetch('/shorten', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ real_url: urlInput }),
  });
  console.log(response);
  const result = await response.json();
  if (response.ok) {
    const shortUrlElement = document.getElementById('short-url');
    shortUrlElement.href = result.shortURL;
    shortUrlElement.textContent = domain + "/" + result.shortURL;
    document.getElementById('result').classList.remove('hidden');
  } else {
    alert('Error shortening URL');
  }
});


