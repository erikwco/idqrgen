window.addEventListener('DOMContentLoaded', () => {
  const form = document.getElementById('qrForm');
  const modal = document.getElementById('modal');
  const qrImage = document.getElementById('qrImage');
  const closeBtn = document.getElementById('closeBtn');
  const downloadBtn = document.getElementById('downloadBtn');

  form.addEventListener('submit', async (e) => {
    e.preventDefault();
    const data = new FormData(form);
    try {
      const resp = await fetch('/generator/', { method: 'POST', body: data });
      if (!resp.ok) throw new Error('Failed to generate QR');
      const blob = await resp.blob();
      const url = URL.createObjectURL(blob);
      qrImage.src = url;
      downloadBtn.href = url;
      modal.classList.remove('hidden');
    } catch (err) {
      console.error(err);
    }
  });

  closeBtn.addEventListener('click', () => {
    modal.classList.add('hidden');
    if (qrImage.src.startsWith('blob:')) {
      URL.revokeObjectURL(qrImage.src);
    }
    qrImage.src = '';
    downloadBtn.removeAttribute('href');
  });
});
