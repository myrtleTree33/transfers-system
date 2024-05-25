"use client"

import Image from 'next/image'
import {API} from "@stoplight/elements";
import docJson  from '@backend/swagger.json';

export default function Home() {

  const swaggerJson = {...docJson}

  return (
<div>
<API
apiDescriptionDocument={swaggerJson}
// basePath={swaggerJson.servers[0].url}
hideInternal={true}
layout={"sidebar"}
/>
</div>
  )
}
