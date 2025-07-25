export function formatDate(upcomingDate) {
  const date = new Date(upcomingDate);
  const month = new Intl.DateTimeFormat('en-US', { month: 'long' }).format(date);
  const day = date.getUTCDate();

  const formattedDate = `${month} ${day}`;
  return formattedDate.toUpperCase();
}

export function formatDateMDY(dateStr) {
  const date = new Date(dateStr);
  const month = String(date.getMonth() + 1);
  const day = String(date.getDate());
  const year = date.getFullYear().toString().slice(-2);
  return `${month}/${day}/${year}`;
}