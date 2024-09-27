document.getElementById('url-form').addEventListener('submit', async function(event) {
  const shortUrlElement = document.getElementById('short-url');
  const resultDiv = document.getElementById('result');
  const errorElement = document.getElementById('error');
  resultDiv.classList.add('hidden');
  errorElement.classList.add('hidden');
  const domainWithoutHttp = window.location.origin.replace('http://', '').replace('https://', '');
  event.preventDefault();
  const urlInput = document.getElementById('url-input').value;
  const response = await fetch('/api/shorten', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ real_url: urlInput }),
  }).catch((error) => {
    console.error('Error:', error);
    errorElement.classList.remove('hidden');
  }).finally(() => {
    console.log('Request completed')
  });
  console.log(response);
  const result = await response.json();
  if (response.ok) {
    errorElement.classList.add('hidden');
    resultDiv.classList.remove('hidden');
    shortUrlElement.href = result.shortURL;
    shortUrlElement.textContent = domainWithoutHttp + "/" + result.shortURL;
    const resultElement = document.getElementById('result');
    resultElement.classList.remove('hidden');
    resultElement.classList.add('flex')
  } else {
    errorElement.classList.remove('hidden');
    errorElement.textContent = result.error;
  }
});


// Copy to clipboard
document.getElementById('copy-button').addEventListener('click', function() {
  const shortUrlElement = document.getElementById('short-url');
  navigator.clipboard.writeText(shortUrlElement.text).then(function() {
    console.log('Copied to clipboard');
    const copyButton = document.getElementById('copy-button');
    // create tooltip and remove after 1 second
    createToolTip(copyButton);
    setTimeout(() => {
      removeToolTip(copyButton);
    }, 1000);
  }, function(err) {
    console.error('Error copying to clipboard', err);
  });
});


function createToolTip(elementToAttachTo) {
  const tooltip = document.createElement('div');
  tooltip.classList.add('tooltip');
  tooltip.textContent = 'Copied!';
  tooltip.style.position = 'absolute';
  tooltip.style.top = '-10px';
  tooltip.style.left = '0';
  tooltip.style.backgroundColor = 'black';
  tooltip.style.color = 'white';
  tooltip.style.padding = '5px';
  tooltip.style.borderRadius = '5px';
  tooltip.style.zIndex = '1000';
  elementToAttachTo.appendChild(tooltip);
}

function removeToolTip(element) {
  const tooltip = element.querySelector('.tooltip');
  if (tooltip) {
    tooltip.remove();
  }
}

