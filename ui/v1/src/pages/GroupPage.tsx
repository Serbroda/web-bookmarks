import LinkCard, { LinkItem } from '@components/LinkCard';
import { useEffect, useState } from 'react';
import { formatUrl, isValidHttpUrl } from '@utils/url.utils';
import TopNav from '@components/TopNav';
import { ArrowRightIcon, EllipsisVerticalIcon } from '@heroicons/react/20/solid';
import LinkModal from '@components/modals/LinkModal';
import useLinkModal from '@stores/modals/useLinkModal';

const mockItems: LinkItem[] = [
  {
    title: 'www.google.de',
    url: 'https://www.google.de',
  },
  {
    title: 'heise',
    url: 'https://heise.de',
  },
  {
    title: 'golem.de',
    url: 'https://golem.de',
  },
  {
    title: 'Entwickler Info',
    url: 'https://dev.to',
  },
];

const GroupPage = () => {
  const [search, setSearch] = useState<string>('');
  const [hasValidUrl, setHasValidUrl] = useState(false);
  const { openModal: openLinkModal } = useLinkModal();

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
    <div>
      <TopNav showMenuButton={true}>
        <div className="h-6 w-px mr-4 bg-gray-900/10 md:hidden" aria-hidden="true"></div>

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

      <div className="p-4 h-full overflow-auto flex flex-col items-center">
        <div className="w-full lg:w-[80%] pb-4 relative">
          <input
            id="username"
            name="username"
            type="search"
            autoComplete="username"
            required
            className={`input w-full ${hasValidUrl ? 'input-success' : ''}`}
            placeholder="Search or add link"
            onChange={(event) => setSearch(event.target.value)}
          />
          {search && hasValidUrl && (
            <button
              type="submit"
              className="btn btn-xs btn-ghost text-gray-400 hover:text-indigo-400 hover:bg-gray-50 absolute right-8 top-1 w-24"
              onClick={() => {
                openLinkModal({
                  mode: 'new',
                  url: search,
                  onSave: () => {
                    console.log('Link saved');
                  },
                });
              }}>
              <span className="pr-2">Add link</span> <ArrowRightIcon className="h-4 w-4" />
            </button>
          )}
        </div>

        <div className="w-full grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
          {mockItems
            .filter((item) => {
              if (!search || !search.trim()) {
                return true;
              }
              const text = search.toLowerCase().trim();
              return item.title.toLowerCase().includes(text) || item.url.toLowerCase().includes(text);
            })
            .map((item, idx) => (
              <LinkCard key={idx} item={item} />
            ))}
        </div>
      </div>

      <LinkModal />
    </div>
  );
};

export default GroupPage;
