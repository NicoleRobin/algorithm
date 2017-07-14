#include <stdio.h>
#include <stdlib.h>

struct listNode
{
	int data;
	struct listNode *next;
};

struct listNode* reveriseList1(struct listNode *head)
{
	struct listNode *p = head;
	struct listNode *q = NULL;
	struct listNode *pRet = NULL;
	while (p != NULL)
	{
		q = p->next;
		p->next = pRet;
		pRet = p;
		p = q;
	}
	return pRet;
}

struct listNode* reveriseList2(struct listNode *head)
{
	if (head == NULL || head->next == NULL)
	{
		return head;
	}
	else
	{
		struct listNode *p = reveriseList2(head->next);
		// 找到尾节点
		struct listNode *tail = p;
		while (tail->next)
		{
			tail = tail->nex;
		}
		tail->next = head;
		head->next = NULL;
		
		return p;
	}
}

int main(void)
{
	struct listNode* head = NULL;
	struct listNode* p = head;
	int n;
	int temp;
	while (scanf("%d", &n) != EOF)
	{
		// build list
		head = NULL;
		for (int i = 0; i < n; i++)
		{
			scanf("%d", &temp);
			struct listNode *pNode = (struct listNode*)malloc(sizeof(struct listNode));
			pNode->data = temp;
			pNode->next = NULL;
			if (p == NULL)
			{
				head = pNode;
				p = head;
			}
			else
			{
				p->next = pNode;
				p = pNode;
			}
		}

		// reverise list
		head = reveriseList2(head);

		// print list
		p = head;
		if (p == NULL)
		{
			printf("NULL");
		}
		else
		{
			struct listNode *q = NULL;
			while (p != NULL)
			{
				printf("%d", p->data);
				if (p->next != NULL)
				{
					printf(" ");
				}
				q = p->next;
				free(p);
				p = q;
			}
		}
		printf("\n");
	}
	return 0;
}
