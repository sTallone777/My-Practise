#include <iostream>

using namespace std;

struct node *insert_beg(struct node *, int);
struct node *insert_end(struct node *, int);
struct node *insert_after(struct node *, int);
struct node *insert_before(struct node *, int);

struct node *create(struct node *);

struct node *del_beg(struct node *);
struct node *del_end(struct node *);
struct node *del_after(struct node *);
struct node *del_before(struct node *);
struct node *del_element(struct node *);

void count(struct node *);
void display(struct node *);

struct node{
    int data;
    node *next;
};

node *insert_beg(node *start, int value){
    node *new_node = new node;
    new_node->data = value;
    new_node->next = start;
    start = new_node;
    return start;
}

node *insert_end(node *start, int value){
    node *tmp = start;
    node *new_node = new node;
    new_node->data = value;
    new_node->next = NULL;

    while(tmp->next != NULL){
        tmp = tmp->next;
    }

    tmp->next = new_node;
    return start;
}

node *insert_before(node *start, int value){
    int num;
    node *tmp = start;
    node *preTmp = tmp;

    cout<<"Enter the number before which to insert: ";
    cin>>num;

    node *new_node = new node;

    new_node->data = value;

    //find pre node
    while(tmp->data != num){
        preTmp = tmp;
        tmp = tmp->next;
    }

    preTmp->next = new_node;
    new_node->next = tmp;
    return start;
}

node *insert_after(node *start, int value){
    int num;
    node *tmp = start;
    node *preTmp = tmp;

    cout<<"Enter the number after which to insert: ";
    cin>>num;

    node *new_node = new node;

    new_node->data = value;

    //find pre node
    while(preTmp->data != num){
        preTmp = tmp;
        tmp = tmp->next;
    }

    preTmp->next = new_node;
    new_node->next = tmp;
    return start;
}

node *create(node *start){
    int value;

    cout<<"Enter -1 to stop\n";
    cout<<"Enter the value to be stored: ";
    cin>> value;

    while(value != -1){
        if(start == NULL){
            node *new_node = new node;
            new_node->next = NULL;
            new_node->data = value;
            start = new_node;
        }else{
            insert_end(start, value);
        }

        cout<<"Enter the value to be stored: ";
        cin>>value;
    }

    return start;
}

node *del_beg(node *start){
    if(start == NULL){
        cout<<"The linked list is empty !\n";
    }else{
        node *ptr = start;
        start = start->next;

        cout<<"The deleted value is :"<<ptr->data<<endl;
        delete ptr;
    }

    return start;
}

node *del_end(node *start){
    if(start == NULL){
        cout<<"The linked list is empty !\n";
    }else{
        node *ptr = start;
        node *preptr = ptr;
        
        while(ptr->next != NULL){
            preptr = ptr;
            ptr = ptr->next;
        }

        preptr->next = NULL;
        cout<<"The deleted value is: "<<ptr->data<<endl;
        delete ptr;
    }
    return start;
}

node *del_after(node *start){
    if(start == NULL){
        cout<<"The linked is empty !\n";
    }else{
        int num;

        cout<<"Enter the element after which to be deleted: ";
        cin>>num;

        node *ptr = start;
        node *preptr = ptr;

        while(preptr->data != num){
            preptr = ptr;
            ptr = ptr->next;
        }

        preptr->next = ptr->next;
        cout<<"The deleted value is: "<<ptr->data<<endl;
        delete ptr;
    }
    return start;
}

node *del_before(node *start){
    if(start == NULL){
        cout<<"The linked is empty !\n";
    }else{
        int num;

        cout<<"Enter the element before which to be deleted: ";
        cin>>num;

        node *ptr = start;
        node *preptr = ptr;

        while(ptr->next->data != num){
            preptr = ptr;
            ptr = ptr->next;
        }

        preptr->next = ptr->next;
        cout<<"The deleted value is: "<<ptr->data<<endl;
        delete ptr;
    }
    return start;
}

node *del_element(node *start){
    if(start == NULL){
        cout<<"The linked list is empty !\n";
    }else{
        int num;
        cout<<"Enter the element which to be deleted: ";
        cin>>num;

        node *ptr = start;
        node *preptr = ptr;

        while(ptr->data != num){
            preptr = ptr;
            ptr = ptr->next;
        }

        preptr->next = ptr->next;
        cout<<"The deleted value is: "<<ptr->data<<endl;
        delete ptr;
    }
    return start;
}

void count(node *start){
    int count =0;

    if(start == NULL){
        cout<<"The linked list is empty !\n";
    }else{
        node *ptr = start;

        while(ptr != NULL){
            ptr = ptr->next;
            count++;
        }
    }
    cout<<"The number of elements present in the linked list are: "
    <<count<<endl;
}

void display(node *start){
    if(start == NULL){
        cout<<"The linked list is empty !\n";
    }else{
        node *ptr = start;
        while(ptr != NULL){
            cout<<ptr->data<<" ";
            ptr = ptr->next;
        } 
        cout<<endl;
    }
}

int main(){
    node *start = NULL;
    int option, value;

    do{
        cout<<"\n************* MENU ***************\n"
        <<"1. Create the linked list\n"
        <<"2. Insert at the beginning\n"
        <<"3. Insert at the end\n"
        <<"4. Insert after an element\n"
        <<"5. Insert before an element\n"
        <<"6. Delete from the beginning\n"
        <<"7. Delete from the end\n"
        <<"8. Delete after an element\n"
        <<"9. Delete before an element\n"
        <<"10. Delete the element\n"
        <<"11. Count\n" 
        <<"12. Display\n"
        <<"13. Exit\n";

        cout<<"Enter you option: ";
        cin>>option;

        switch(option){
            case 1:
                start = create(start);
                break;
            case 2:
                cout<<"Enter the value to be stored: ";
                cin>>value;
                start = insert_beg(start, value);
                break;

            case 3:
                cout<<"Enter the value to be stored: ";
                cin>>value;
                start = insert_end(start, value);
                break;
            case 4:
                cout<<"Enter the value to be stored: ";
                cin>>value;
                start = insert_after(start, value);
                break;
            case 5:
                cout<<"Enter the value to be stored: ";
                cin>>value;
                start = insert_before(start, value);
                break;
            case 6:
                start = del_beg(start);
                break;
            case 7:
                start = del_end(start);
                break;
            case 8:
                start = del_after(start);
                break;
            case 9:
                start = del_before(start);
                break;
            case 10:
                start = del_element(start);
                break;
            case 11:
                count(start);
                break;
            case 12:
                display(start);
                break;
            case 13:
                break;
            default:
                cout<<"Wrong option is selected !!\n";
                break;
        }
    }while(option != 13);

    if(option == 13){
        cout<<"\nThank you for using the program !\n"
        <<"Have a good day.\n\n";
        delete start;
    }
}