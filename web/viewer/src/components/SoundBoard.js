import select from '../audio/select.wav';
import blip1 from '../audio/blip1.wav';
import blip2 from '../audio/blip2.wav';

const SoundBoard = {
  Select: () => {
    return <audio src={select} autoPlay />
  },
  Blip1: () => {
    return <audio src={blip1} autoPlay />
  },
  Blip2: () => {
    return <audio src={blip2} autoPlay />
  }
}

export default SoundBoard;
