import BasicGraph from './BasicGraph';
import Example from './example';

export default function Home() {

  return (
      <div>
        <h2>test</h2>
        <Example width={600} height={300} />

        <BasicGraph width={800} height={600} />
      </div>
  );
}
