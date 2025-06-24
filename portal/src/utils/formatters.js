export function formatDate(upcomingDate) {
  const date = new Date(upcomingDate);
  const month = new Intl.DateTimeFormat('en-US', { month: 'long' }).format(date);
  const day = date.getUTCDate();

  const formattedDate = `${month} ${day}`;
  return formattedDate.toUpperCase();
}