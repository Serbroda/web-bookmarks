import LinkCard from '@components/LinkCard';
import { useEffect, useState } from 'react';
import { formatUrl, isValidHttpUrl } from '@utils/url.utils';
import TopNav from '@components/TopNav';
import { EllipsisVerticalIcon } from '@heroicons/react/20/solid';

const GroupPage = () => {
  const [search, setSearch] = useState<string>('');
  const [hasValidUrl, setHasValidUrl] = useState(false);

  useEffect(() => {
    setHasValidUrl(false);
    let formattedUrl: URL | undefined;

    if (search !== '') {
      formattedUrl = formatUrl(search);
    }

    if (isValidHttpUrl(formattedUrl?.href)) {
      setHasValidUrl(true);
    }
  }, [search]);

  return (
    <>
      <TopNav showMenuButton={true}>
        <div className="flex flex-1 items-stretch justify-start">
          <div className="block">
            <div className="flex space-x-4">Title</div>
          </div>
        </div>
        <div className="absolute inset-y-0 right-0 flex items-center pr-2 sm:static sm:inset-auto sm:ml-6 sm:pr-0">
          <button type="button" className="btn btn-ghost rounded-full ">
            <EllipsisVerticalIcon className="h-5 w-5" />
          </button>
        </div>
      </TopNav>
      <LinkCard
        item={{
          title: 'www.test.de',
          url: 'www.test.de',
        }}
        onClick={(item) => console.log(item)}
      />
    </>
  );
};

export default GroupPage;
