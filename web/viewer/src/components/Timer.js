const Timer = ({timer, status, roundTime}) => {
  if(!timer || status !== 'running') {
    return null;
  } else {
    let time = roundTime - timer;
    const minutes = parseInt(time / 60);
    return (<div className="timer">{minutes}:{time - minutes * 60}</div>);
  }
}

export default Timer;
