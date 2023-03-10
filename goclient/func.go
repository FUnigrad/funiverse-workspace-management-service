package goclient

// func (client *GoClient) CreateWorkspace(workspace model.Workspace) error {
// 	_, err := client.CreateNamespace(workspace.Name)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (client *GoClient) CreateNamespace(name string) (*v1.Namespace, error) {
// 	clientset := client.Client
// 	namespace := &v1.Namespace{
// 		ObjectMeta: metav1.ObjectMeta{
// 			Name: name,
// 		},
// 	}
// 	namespace, err := clientset.CoreV1().Namespaces().Create(context.TODO(), namespace, metav1.CreateOptions{})
// 	return namespace, err
// }
