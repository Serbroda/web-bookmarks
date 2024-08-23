package de.serbroda.ragbag.services;

import de.serbroda.ragbag.models.Account;
import de.serbroda.ragbag.models.Page;
import de.serbroda.ragbag.models.PageAccount;
import de.serbroda.ragbag.models.Space;
import de.serbroda.ragbag.models.shared.PageRole;
import de.serbroda.ragbag.repositories.PageAccountRepository;
import de.serbroda.ragbag.repositories.PageRepository;
import org.springframework.stereotype.Service;

@Service
public class PageService {

    private final PageRepository pageRepository;
    private final PageAccountRepository pageAccountRepository;

    public PageService(PageRepository pageRepository, PageAccountRepository pageAccountRepository) {
        this.pageRepository = pageRepository;
        this.pageAccountRepository = pageAccountRepository;
    }

    public Page createPage(Space space, Page page, Account account) {
        page.setSpace(space);
        page = pageRepository.save(page);
        addAccountToPage(page, account, PageRole.OWNER);
        return page;
    }

    public PageAccount addAccountToPage(Page page, Account account, PageRole pageRole) {
        PageAccount pageAccount = new PageAccount();
        pageAccount.setPage(page);
        pageAccount.setAccount(account);
        pageAccount.setRole(pageRole);
        return pageAccountRepository.save(pageAccount);
    }
}
